import { useEffect, useState } from "react";
import { ncweClient } from "./nwce-client";
import { LogInInfoPopUpText } from "./popup";

export const LogIn = ({ findUser }: { findUser: Function }) => {
  const [Username, setUsername] = useState<string>("");
  const [Password, setPassword] = useState<string>("");
  const [errorText, setErrorText] = useState<string>("");

  const handleLogIn = async () => {
    const lio = await ncweClient.LogIn({ Username, Password });
    if (lio.error) {
      setErrorText(lio.error);
      return;
    }
    ncweClient.setSession(lio.output.Session);
    localStorage.setItem("session", lio.output.Session);
    await findUser();
  };

  useEffect(() => {
    const handleKeyPress = (event: KeyboardEvent) => {
      if (event.key === "Enter") handleLogIn();
    };
    document.addEventListener("keydown", handleKeyPress);
    return () => document.removeEventListener("keydown", handleKeyPress);
  }, [Username, Password]);

  return (
    <div style={{ display: "flex", flexDirection: "column", alignItems: "center", width: "100vw", height: "100vh" }}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          marginTop: "100px",
          padding: "20px 40px",
          backgroundColor: "#252525",
          border: "2px solid #3a3a3a",
          boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px",
          borderRadius: "8px",
          gap: "8px",
        }}
      >
        <label style={{ color: "white", fontFamily: "monospace" }}>nginx config web editor</label>
        <input placeholder="username" onChange={(e) => setUsername(e.target.value)} value={Username} />
        <input placeholder="password" type="password" onChange={(e) => setPassword(e.target.value)} value={Password} />
        <label style={{ color: "white", fontFamily: "monospace" }}>&zwnj;{errorText}&zwnj;</label>
        <div style={{ display: "flex", gap: "10px" }}>
          <button style={{ minWidth: "75px" }} onClick={() => alert(LogInInfoPopUpText)}>
            info
          </button>
          <button style={{ minWidth: "75px" }} onClick={handleLogIn}>
            log in
          </button>
        </div>
      </div>
    </div>
  );
};
