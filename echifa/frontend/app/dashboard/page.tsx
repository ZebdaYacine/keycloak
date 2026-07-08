import { auth } from "@/auth";
import { Header } from "@/app/_lib/components/Header";
import { redirect } from "next/navigation";

async function getCards(accessToken: string) {
  const res = await fetch("http://167.86.79.16:8902/api/admin/cards", {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
    cache: "no-store",
  });

  if (!res.ok) {
    throw new Error("Failed to load cards");
  }

  return res.json();
}

export default async function DashboardPage() {
  const session = await auth();

  if (!session?.user || !session.accessToken) {
    redirect("/");
  }

  const cards = await getCards(session.accessToken);

  return (
    <>
      <Header session={session} />

      <main className="mx-auto max-w-5xl p-8">
        <section className="rounded-xl border bg-white p-8">
          <h1 className="text-3xl font-bold">Protected Dashboard</h1>

          <p className="mt-6 text-slate-600">
            You are authenticated with Keycloak.
          </p>

          <h2 className="mt-8 text-xl font-semibold">User</h2>
          <pre className="mt-4 rounded-xl border p-4">
            {JSON.stringify(session.user, null, 2)}
          </pre>

          <h2 className="mt-8 text-xl font-semibold">Cards</h2>
          <pre className="mt-4 rounded-xl border p-4">
            {JSON.stringify(cards, null, 2)}
          </pre>
        </section>
      </main>
    </>
  );
}
