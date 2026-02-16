// BARK: Remove console.log statements before production
console.log("Debug mode enabled");

function calculateSum(a, b) {
  // Regular comment
  // BARK
  return a + b;
}

/* BARK: This is a temporary workaround
   Need to implement proper solution */
function temporaryFix() {
  // BARK: Delete this entire function later
  // BARK - plain marker with dash
  return "temp";
}
