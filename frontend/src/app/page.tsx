import { AuthContainer } from "../container/authContainer";

export default function Home() {
  return (
    <div className="font-sans">
      <main className="bg-white">
        <div className="max-w-2xl mx-auto py-32 px-16">
          <AuthContainer />
        </div>
      </main>
    </div>
  );
}
