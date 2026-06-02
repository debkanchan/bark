#!/usr/bin/env node
// Prepare the npm packages for publishing.
//
//   node npm/build.mjs <version> <artifacts-dir>
//
// - <version>      e.g. "0.1.2" (no leading "v")
// - <artifacts-dir> directory containing the per-platform release binaries,
//                   as produced by the Release workflow's build job and
//                   downloaded via actions/download-artifact. Expected layout:
//                     <artifacts-dir>/bark-<goos>-<goarch>/bark-<goos>-<goarch>[.exe]
//
// The script:
//   1. stamps every package.json (main + platform pkgs + optionalDependencies)
//      with <version>, and
//   2. copies each platform binary into its package's bin/ dir (chmod +x).
//
// Run from the repo root.

import { promises as fs } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const NPM_DIR = path.dirname(fileURLToPath(import.meta.url));

// goos/goarch (release artifact naming) -> npm package directory + binary name
const PLATFORMS = [
  { goos: "linux", goarch: "amd64", dir: "bark-linux-x64", bin: "bark" },
  { goos: "linux", goarch: "arm64", dir: "bark-linux-arm64", bin: "bark" },
  { goos: "darwin", goarch: "amd64", dir: "bark-darwin-x64", bin: "bark" },
  { goos: "darwin", goarch: "arm64", dir: "bark-darwin-arm64", bin: "bark" },
  { goos: "windows", goarch: "amd64", dir: "bark-win32-x64", bin: "bark.exe" },
];

async function readJson(file) {
  return JSON.parse(await fs.readFile(file, "utf8"));
}

async function writeJson(file, obj) {
  await fs.writeFile(file, JSON.stringify(obj, null, 2) + "\n");
}

async function setVersion(version) {
  // Main package: own version + every optionalDependency version.
  const mainPath = path.join(NPM_DIR, "bark", "package.json");
  const main = await readJson(mainPath);
  main.version = version;
  for (const dep of Object.keys(main.optionalDependencies ?? {})) {
    main.optionalDependencies[dep] = version;
  }
  await writeJson(mainPath, main);
  console.log(`stamped ${main.name}@${version}`);

  // Platform packages.
  for (const p of PLATFORMS) {
    const pkgPath = path.join(NPM_DIR, p.dir, "package.json");
    const pkg = await readJson(pkgPath);
    pkg.version = version;
    await writeJson(pkgPath, pkg);
    console.log(`stamped ${pkg.name}@${version}`);
  }
}

async function copyBinaries(artifactsDir) {
  for (const p of PLATFORMS) {
    const srcName = `bark-${p.goos}-${p.goarch}${p.goos === "windows" ? ".exe" : ""}`;
    const src = path.join(artifactsDir, `bark-${p.goos}-${p.goarch}`, srcName);
    const destDir = path.join(NPM_DIR, p.dir, "bin");
    const dest = path.join(destDir, p.bin);

    await fs.mkdir(destDir, { recursive: true });
    await fs.copyFile(src, dest);
    await fs.chmod(dest, 0o755);
    console.log(`copied ${src} -> ${dest}`);
  }
}

async function main() {
  const [, , version, artifactsDir] = process.argv;
  if (!version) {
    console.error("usage: node npm/build.mjs <version> [artifacts-dir]");
    process.exit(1);
  }
  await setVersion(version);
  if (artifactsDir) {
    await copyBinaries(artifactsDir);
  } else {
    console.log("no artifacts dir given; skipped binary copy (version-only run)");
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
