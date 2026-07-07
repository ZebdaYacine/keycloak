import type { Session } from "next-auth";
import Link from "next/link";
import { SignInButton } from "./SignInButton";
import { SignOutButton } from "./SignOutButton";

export function Header({ session }: { session: Session | null }) {
  return (
    <header className="border-b bg-white">
      <div className="mx-auto flex max-w-5xl items-center justify-between p-4">
        <Link className="text-xl font-bold" href="/">Next.js + Keycloak</Link>
        <nav className="flex items-center gap-4">
          <Link className="text-sm text-slate-700 hover:underline" href="/dashboard">Dashboard</Link>
          {session?.user ? (
            <div className="flex items-center gap-3">
              <span className="text-sm text-slate-700">{session.user.name ?? session.user.email}</span>
              <SignOutButton />
            </div>
          ) : (
            <SignInButton />
          )}
        </nav>
      </div>
    </header>
  );
}
