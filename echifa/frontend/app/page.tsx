import { auth } from "@/auth";
import { Header } from "@/app/_lib/components/Header";
import { SignInButton } from "@/app/_lib/components/SignInButton";
import Link from "next/link";
import Image from "next/image";

export default async function HomePage() {
  const session = await auth();

  return (
    <>
      <Header session={session} />

      <main className="mx-auto flex min-h-[calc(100vh-80px)] max-w-7xl items-center px-8 pt-16 pb-8">
        {" "}
        <section className="flex w-full flex-col items-center justify-center gap-16 md:flex-row md:gap-24">
          {/* Left - Logo */}
          <div className="flex w-full justify-center md:w-1/2">
            <Image
              src="/cnas.png"
              alt="CNAS Logo"
              width={200}
              height={200}
              priority
              className="h-auto max-w-full"
            />
          </div>

          {/* Right - Content */}
          <div className="w-full md:w-1/2">
            <h1 className="text-4xl font-bold text-slate-900">
              ECHIFA CNAS Digital Card Platform
            </h1>

            <p className="mt-6 text-lg leading-8 text-slate-600">
              Secure ECHIFA application for CNAS users. The platform uses
              Keycloak SSO to authenticate users and allows authorized access to
              digital CHIFA card information.
            </p>

            <div className="mt-8">
              {session?.user ? (
                <Link
                  href="/dashboard"
                  className="inline-block rounded-lg bg-blue-600 px-6 py-3 font-medium text-white transition hover:bg-blue-700"
                >
                  Open ECHIFA Dashboard
                </Link>
              ) : (
                <SignInButton />
              )}
            </div>
          </div>
        </section>
      </main>
    </>
  );
}
