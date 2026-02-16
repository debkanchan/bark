// Clean TypeScript file
interface User {
  id: number;
  name: string;
  // Regular field comment
  email: string;
}

/* Implementation of user service */
function getUser(id: number): User {
  // Standard implementation
  return {
    id: id,
    name: "User",
    email: "user@example.com",
  };
}

// Configuration object
const config = {
  apiUrl: "https://api.example.com",
  timeout: 5000,
};
