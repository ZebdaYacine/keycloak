import { auth } from "@/auth";
import { Header } from "@/app/_lib/components/Header";
import { SignInButton } from "@/app/_lib/components/SignInButton";
import Link from "next/link";

export default async function HomePage() {
  const session = await auth();

  return (
    <>
      <Header session={session} />

      <main className="mx-auto max-w-5xl p-8">
        <section className="rounded-xl border bg-white p-8">
          <h1 className="text-3xl font-bold">
            SEGMA CNAS Secure Document Platform
          </h1>

          <p className="mt-6 text-slate-600">
            Secure SEGMA application for CNAS users. The platform uses Keycloak
            SSO to authenticate users and allows authorized access to sickness
            proof upload and document management services.
          </p>

          <div className="mt-6 flex gap-4">
            {session?.user ? (
              <Link
                className="rounded-md bg-blue-600 px-4 py-2 font-medium text-white"
                href="/dashboard"
              >
                Open SEGMA Dashboard
              </Link>
            ) : (
              <SignInButton />
            )}
          </div>
        </section>
      </main>
    </>
  );
}
