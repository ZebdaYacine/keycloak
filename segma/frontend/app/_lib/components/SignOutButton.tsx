"use client";

import { signOut } from "@/app/_lib/actions/auth";

export function SignOutButton() {
  async function handleLogout() {
    localStorage.clear();
    sessionStorage.clear();
    await signOut();
  }

  return (
    <button
      type="button"
      onClick={handleLogout}
      className="rounded-md bg-slate-800 px-4 py-2 font-medium text-white hover:bg-slate-900"
    >
      Sign out
    </button>
  );
}
