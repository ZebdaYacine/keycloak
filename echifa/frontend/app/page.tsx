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

      <Image
        src="../cnas.png"
        width={500}
        height={500}
        alt="Picture of the author"
      />

      <main className="mx-auto max-w-5xl p-8">
        <section className="rounded-xl border bg-white p-8">
          <h1 className="text-3xl font-bold">
            ECHIFA CNAS Digital Card Platform
          </h1>

          <p className="mt-6 text-slate-600">
            Secure ECHIFA application for CNAS users. The platform uses Keycloak
            SSO to authenticate users and allows authorized access to digital
            CHIFA card information.
          </p>

          <div className="mt-6 flex gap-4">
            {session?.user ? (
              <Link
                className="rounded-md bg-blue-600 px-4 py-2 font-medium text-white"
                href="/dashboard"
              >
                Open ECHIFA Dashboard
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
