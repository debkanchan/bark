# Using Bark with GitHub Desktop

GitHub Desktop is a popular Git client, but it has historically had issues with git hooks. This guide explains how to use Bark effectively with GitHub Desktop.

## The Problem

GitHub Desktop **does not run pre-push hooks** by default. This means the standard `bark git-hook install` command won't work with GitHub Desktop - you'll be able to push code with BARK comments without any warnings.

## The Solution: Pre-Commit Hooks

**Good news!** GitHub Desktop **does support pre-commit hooks**. We've added a new command to Bark that installs a pre-commit hook instead of a pre-push hook.

### Installation

```bash
# Install bark
go install github.com/debkanchan/bark/cmd/bark@latest

# Install the pre-commit hook (works with GitHub Desktop!)
cd /path/to/your/repo
bark git-hook install-commit
```

### How It Works

With the pre-commit hook installed:

1. You add BARK comments to your code as reminders
2. When you try to commit in GitHub Desktop, Bark runs automatically
3. If BARK comments are found, the commit is **blocked**
4. You'll see an error message in GitHub Desktop showing where the BARK comments are
5. Remove the BARK comments and try again

### Example

Let's say you have this code with BARK comments:

```typescript
// BARK: Remove this debug code
console.log("Debug mode enabled");

// BARK: Replace with production API key
const apiKey = "test-123";
```

When you try to commit in GitHub Desktop:

```
🐕 Running bark to check for BARK comments...
Found 2 BARK comment(s):

src/components/Backgrounds.tsx:11:1: // BARK: Remove this debug code
src/components/Backgrounds.tsx:14:1: // BARK: Replace with production API key

❌ Commit blocked: BARK comments found
Please remove BARK comments before committing
Run 'bark .' to see all BARK comments
```

The commit will be rejected! GitHub Desktop will show this error in the commit panel.

## Uninstalling

If you want to remove the pre-commit hook:

```bash
bark git-hook uninstall-commit
```

## Can I Use Both Pre-Commit and Pre-Push?

Yes! You can install both hooks for double protection:

```bash
# Install pre-commit (for GitHub Desktop and early catching)
bark git-hook install-commit

# Install pre-push (for CLI git users as backup)
bark git-hook install
```

## Troubleshooting

### The hook isn't running

1. Make sure you're in the repository root when installing
2. Check that `.git/hooks/pre-commit` exists and is executable:
   ```bash
   ls -la .git/hooks/pre-commit
   ```
3. Verify bark is in your PATH:
   ```bash
   which bark
   # Should show: /Users/yourusername/go/bin/bark
   ```

### GitHub Desktop shows "hook failed" with no details

This usually means `bark` isn't in GitHub Desktop's PATH. Try:

1. Make sure bark is installed: `which bark`
2. Add Go's bin directory to your PATH in `~/.zshrc` or `~/.bash_profile`:
   ```bash
   export PATH="$HOME/go/bin:$PATH"
   ```
3. Restart GitHub Desktop

### I want to commit anyway (override the hook)

If you absolutely need to bypass the hook temporarily:

```bash
git commit --no-verify -m "your message"
```

**Warning:** This defeats the purpose of Bark. Only use this if you really know what you're doing!

## Why This Approach?

We chose to create a separate `install-commit` command instead of making `install` automatically detect GitHub Desktop because:

1. **User choice**: Some developers prefer catching BARK comments at push time (allows WIP commits)
2. **Explicit behavior**: It's clearer what each command does
3. **Compatibility**: You can install both hooks if needed

## Resources

- [GitHub Desktop Documentation](https://docs.github.com/en/desktop)
- [Git Hooks Documentation](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [Bark Main README](README.md)
