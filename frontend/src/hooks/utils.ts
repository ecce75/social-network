export async function useAuth() {
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
  try {
      const res = await fetch(`${FE_URL}:${BE_PORT}/api/users/check-auth`, {
          credentials: 'include'
      }), data = await res.json();
      return data; // Return the data received from the fetch call
  } catch (error) {
      return { isAuthenticated: false }; // Return a default value in case of error
  }
}

export function formatDate(string: string) {
    var options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: 'numeric'};
    return new Date(string).toLocaleDateString([], options);
}
