import Link from "next/link";
import { Button } from "../ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import { Input } from "../ui/input";
import PersonIcon from '@mui/icons-material/Person';
import PasswordIcon from '@mui/icons-material/Password';

interface AuthProps {
    title: string;
    description: string;
    username: string;
    password: string;
    setUsername: (username: string) => void;
    setPassword: (password: string) => void;
    handleClick: () => void;
}

export const Auth = ({ title, description, username, password, setUsername, setPassword, handleClick }: AuthProps) => {
    return (
        <div>
        <Card>
            <CardHeader>
                <CardTitle>
                    {title}
                </CardTitle>
                <CardDescription>
                    {description}
                </CardDescription>
            </CardHeader>
            <CardContent>
                    <form action="/login" method="post">
                        <div className="flex flex-col gap-4 py-4">
                            <PersonIcon className="margin-top-2"/>
                            <Input type="email" placeholder="Email" name="email" value={username} onChange={(e) => setUsername(e.target.value)} />
                            <PasswordIcon className="margin-top-2" />
                            <Input type="password" placeholder="Password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                        </div>
                        <div className="flex justify-center items-center gap-4 py-4">
                            <Button 
                            className="w-1/2" 
                            type="submit"
                            onClick={handleClick}
                            >ログイン</Button>
                            <Link href="/signup" className="w-1/2 text-center border border-gray-300 rounded-md p-2">新規登録</Link>
                        </div>
                    </form>
                </CardContent>
        </Card>
        </div>
    );
}