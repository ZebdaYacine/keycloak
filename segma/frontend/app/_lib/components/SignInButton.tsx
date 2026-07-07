import { signIn } from "@/app/_lib/actions/auth";

export function SignInButton() {
  return (
    <form action={signIn}>
      <button className="rounded-md bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700" type="submit">
        Sign in with Keycloak
      </button>
    </form>
  );
}
