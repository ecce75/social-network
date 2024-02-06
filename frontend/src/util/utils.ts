export async function useAuth() {
  try {
      const res = await fetch('http://localhost:8080/api/users/check-auth', {
        credentials: 'include'
      });
      const data = await res.json();
      return data; // Return the data received from the fetch call
  } catch (error) {
      return { isAuthenticated: false }; // Return a default value in case of error
  }
}
