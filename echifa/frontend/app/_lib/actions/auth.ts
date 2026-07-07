"use server";

import { signIn as rawSignIn, signOut as rawSignOut } from "@/auth";

export async function signIn() {
  await rawSignIn(
    "keycloak",
    {
      redirectTo: "/dashboard",
    },
    {
      prompt: "login",
    },
  );
}

export async function signOut() {
  await rawSignOut({
    redirectTo: "/",
  });
}
