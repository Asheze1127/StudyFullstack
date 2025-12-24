"use client";
import { Auth } from "../component/auth/auth";
import { useState } from "react";

export default function Home() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="font-sans">
      <main className="bg-white">
        <div className="max-w-2xl mx-auto py-32 px-16">
          <Auth 
          title="ログイン" 
          description="ログインしてください" 
          username={username} 
          password={password} 
          setUsername={setUsername} 
          setPassword={setPassword} 
          />
        </div>
      </main>
    </div>
  );
}
