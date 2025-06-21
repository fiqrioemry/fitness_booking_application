import { useEffect } from "react";

const GoogleLogin = () => {
  useEffect(() => {
    window.google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: handleCredentialResponse,
    });

    window.google.accounts.id.renderButton(
      document.getElementById("google-signin"),
      { theme: "outline", size: "large" }
    );
  }, []);

  const handleCredentialResponse = async (response) => {
    const idToken = response.credential;
    try {
      const res = await fetch("http://localhost:5000/api/auth/google-signin", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ idToken }),
      });
      const data = await res.json();
      console.log("Login Success", data);
    } catch (err) {
      console.error("Login Failed", err);
    }
  };

  return <div id="google-signin"></div>;
};

export default GoogleLogin;
