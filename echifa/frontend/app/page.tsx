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

      <main className="mx-auto max-w-6xl p-8">
        <section className="grid grid-cols-1 items-center gap-12 rounded-xl border bg-white p-8 md:grid-cols-2">
          {/* Logo */}
          <div className="flex justify-center">
            <Image
              src="/cnas.png"
              width={420}
              height={420}
              alt="CNAS Logo"
              priority
              className="h-auto w-full max-w-[420px]"
            />
          </div>

          {/* Content */}
          <div>
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
                  className="rounded-md bg-blue-600 px-6 py-3 font-medium text-white hover:bg-blue-700 transition-colors"
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
