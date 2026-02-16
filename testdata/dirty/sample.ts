// BARK: Remove this test interface before production
// BARK
interface User {
  id: number;
  name: string;
  // BARK: Add proper validation
  email: string;
}

/* BARK: This is a temporary mock implementation
   Replace with actual API call */
function getUser(id: number): User {
  // BARK: Remove hardcoded data
  return {
    id: 1,
    name: "Test User",
    email: "test@example.com",
  };
}

// Regular comment
const config = {
  // BARK: Use environment variables
  apiUrl: "http://localhost:3000",
  // BARK plain marker test
  timeout: 5000,
};
