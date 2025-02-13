import { redirect } from "next/navigation";

export default function Home() {
  redirect("/auth/login"); // Redirige vers la page de connexion
  return null;
}
