import { Button } from "../ui/button"

type HeaderProps = {
    action: string
}

export default function Header({ action }: HeaderProps) {
    return (
        <div className="sticky top-0 z-50 bg-white dark:bg-black w-full">
            <header className="flex justify-between items-center p-4">
                <div className="text-2xl font-bold">SNS</div>
                <Button variant="outline">{action}</Button>
            </header>
        </div>
    );
}