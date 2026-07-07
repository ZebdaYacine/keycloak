import NextAuth, { DefaultSession } from "next-auth";
import { JWT } from "next-auth/jwt";

declare module "next-auth" {
  interface Session {
    accessToken: string;
    idToken: string;
    user: DefaultSession["user"];
    role: string;
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    accessToken: string;
    idToken: string;
    role: string;
  }
}
