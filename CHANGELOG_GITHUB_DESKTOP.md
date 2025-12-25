# Changelog: GitHub Desktop Support

## Summary

Added support for GitHub Desktop and other Git clients by implementing pre-commit hook installation alongside the existing pre-push hook functionality.

## Problem

GitHub Desktop does not execute pre-push hooks, making the original `bark git-hook install` command ineffective for GitHub Desktop users. Users could commit and push code with BARK comments without any warnings.

## Solution

Implemented new commands to install/uninstall pre-commit hooks, which GitHub Desktop does respect.

## Changes Made

### 1. Code Changes (`cmd/bark/main.go`)

#### New Commands
- `bark git-hook install-commit` - Install pre-commit hook
- `bark git-hook uninstall-commit` - Uninstall pre-commit hook

#### New Constants
- `hookContentCommit` - Pre-commit hook script template (similar to pre-push but with different messaging)

#### New Functions
- `installGitHookCommit()` - Installs bark as a pre-commit hook
- `uninstallGitHookCommit()` - Removes bark pre-commit hook

#### Updated Help Text
- Added new commands to usage documentation
- Updated examples to show both hook types

### 2. Documentation Updates

#### Updated `README.md`
- **Features section**: Updated to mention both pre-commit and pre-push hooks with GitHub Desktop support
- **Installation section**: Split into two subsections:
  - "For GitHub Desktop Users (Pre-Commit Hook)"
  - "For CLI Git Users (Pre-Push Hook)"
- **Added explanation**: Why two options exist and when to use each
- **Git Hook Commands section**: Completely restructured to show both hook types with clear examples

#### New `GITHUB_DESKTOP.md`
- Comprehensive guide for GitHub Desktop users
- Explains the problem and solution
- Step-by-step installation instructions
- Troubleshooting section
- Examples of how it works in practice

### 3. Behavior

#### Pre-Commit Hook (`install-commit`)
- Runs before each commit
- Blocks commits if BARK comments are found
- Works with: GitHub Desktop, VS Code, CLI git, and all other Git clients
- Message: "❌ Commit blocked: BARK comments found"

#### Pre-Push Hook (`install`)
- Runs before each push
- Blocks pushes if BARK comments are found
- Works with: CLI git only (bypassed by GitHub Desktop)
- Message: "❌ Push blocked: BARK comments found"

#### Both Hooks Can Coexist
- Users can install both hooks simultaneously
- Provides double protection
- Each hook uses the same marker system (`# BEGIN bark hook` / `# END bark hook`)

## Testing

Tested in Demo-inator repository:
- ✅ Pre-commit hook installed successfully
- ✅ Pre-push hook installed successfully
- ✅ Both hooks coexist in `.git/hooks/` directory
- ✅ Pre-commit hook detects BARK comments
- ✅ Proper error messages displayed

## Backward Compatibility

This change is **100% backward compatible**:
- Existing `bark git-hook install` command works exactly as before
- Existing `bark git-hook uninstall` command unchanged
- No breaking changes to existing functionality
- New commands are purely additive

## Files Modified

1. `cmd/bark/main.go` - Added new commands and functions
2. `README.md` - Updated documentation
3. `GITHUB_DESKTOP.md` - New file (comprehensive guide)

## Migration Guide

### For Existing Users

No action required! Your existing pre-push hooks continue to work.

### For GitHub Desktop Users

Run this command in your repository:

```bash
bark git-hook install-commit
```

You can keep your existing pre-push hook or remove it:

```bash
# Optional: remove pre-push hook if you only use GitHub Desktop
bark git-hook uninstall
```

## Future Considerations

### Potential Enhancements
1. Add `--all` flag to install both hooks at once: `bark git-hook install --all`
2. Auto-detect Git client and suggest appropriate hook
3. Add configuration file support for default hook type
4. Create pre-commit framework integration

### Known Limitations
1. Users must have `bark` in their PATH for hooks to work
2. GitHub Desktop's environment may differ from terminal (rare edge cases)
3. No automatic hook detection - users must choose which to install

## References

- [GitHub Desktop Issue #13112](https://github.com/desktop/desktop/issues/13112) - Git hook support discussion
- [GitHub Community Discussion #24072](https://github.com/orgs/community/discussions/24072) - Pre-commit hooks on Mac
- [Git Hooks Documentation](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
