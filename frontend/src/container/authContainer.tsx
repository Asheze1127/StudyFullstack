"use client";
import { Auth } from "../component/auth/auth"
import { useState } from "react";

export const AuthContainer = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleClick = () => {
        console.log(username, password);
    }

    return (
        <Auth 
            title="ログイン" 
            description="ログインしてください" 
            username={username} 
            password={password} 
            setUsername={setUsername} 
            setPassword={setPassword} 
            handleClick={handleClick}
        />
    )
}