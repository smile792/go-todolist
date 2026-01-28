import { Button } from "@mantine/core";
import { useState } from "react";

export const App = () => {
  const [fetchUser, setFetchUser] = useState<{
    message: string;
    name: string;
    status: number;
    age: number;
  }>();
  const getName = async () => {
    const response = await fetch("http://localhost:8080/user", {
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    setFetchUser(data);
  };
  console.log("fetchName", fetchUser);
  return (
    <>
      <Button onClick={getName}>кликни на меня</Button>
      <div>Имя пользователя {fetchUser?.name}</div>
      <div>Возраст {fetchUser?.age}</div>
    </>
  );
};
