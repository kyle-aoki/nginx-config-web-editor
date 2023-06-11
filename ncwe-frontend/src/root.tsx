import { useEffect, useState } from "react";
import { App } from "./app";
import { LogIn } from "./log-in";
import { ncweClient } from "./nwce-client";
import { GlobalStyles } from "./styled-components";

export default function Root() {
  const [user, setUser] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(true);

  const findUser = async () => {
    const ufo = await ncweClient.UserFind();
    setLoading(false);
    if (ufo.error) return;
    setUser(ufo.output.Username);
  };

  useEffect(() => {
    findUser();
  }, []);

  return (
    <>
      <GlobalStyles />
      {loading ? <div /> : user ? <App /> : <LogIn  findUser={findUser} />}
    </>
  );
}
