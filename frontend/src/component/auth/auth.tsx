import Link from "next/link";
import { Button } from "../ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import { Input } from "../ui/input";
import PersonIcon from '@mui/icons-material/Person';
import PasswordIcon from '@mui/icons-material/Password';
import { useState } from "react";

interface AuthProps {
    username: string;
    password: string;
    setUsername: (username: string) => void;
    setPassword: (password: string) => void;
    handleClick: () => void;
}

export const Auth = ({ username, password, setUsername, setPassword, handleClick }: AuthProps) => {
    const [mode, setMode] = useState<"login" | "signup">("login");

    const handleModeChange = () => {
        setMode(mode === "login" ? "signup" : "login");
    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    {mode === "login" ? "ログイン" : "新規登録"}
                </CardTitle>
                <CardDescription>
                    {mode === "login" ? "ログインしてください" : "新規登録してください"}
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form method="post">
                    <div className="flex flex-col gap-4 py-4">
                        <PersonIcon className="margin-top-2" />
                        <Input type="email" placeholder="Email" name="email" value={username} onChange={(e) => setUsername(e.target.value)} />
                        <PasswordIcon className="margin-top-2" />
                        <Input type="password" placeholder="Password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                    </div>
                    <div className="flex justify-center items-center gap-4 py-4">
                        <Button
                            className="w-1/2"
                            type="submit"
                            onClick={handleClick}
                        >{mode === "login" ? "ログイン" : "新規登録"}</Button>
                        <Button className="w-1/2" type="button" variant="outline" onClick={handleModeChange}>{mode === "login" ? "新規登録" : "ログイン"}</Button>
                    </div>
                </form>
            </CardContent>
        </Card>
    );
}