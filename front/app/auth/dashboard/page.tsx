"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import ThemeToggle from "@/components/ThemeToggle";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { LucideFile as FileIcon } from "lucide-react";

export default function Dashboard() {
  const [user, setUser] = useState<string | null>(null);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadMessage, setUploadMessage] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    fetch("/api/auth/me")
      .then((res) => res.json())
      .then((data) => {
        if (!data.authenticated) {
          router.push("/auth/login");
        } else {
          setUser(data.email);
        }
      });
  }, [router]);

  const handleLogout = async () => {
    await fetch("/api/auth/logout", { method: "POST" });
    router.push("/auth/login");
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setSelectedFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      setUploadMessage("Veuillez sélectionner un fichier.");
      return;
    }

    const formData = new FormData();
    formData.append("file", selectedFile);

    const response = await fetch("/api/upload", {
      method: "POST",
      body: formData,
    });

    if (response.ok) {
      setUploadMessage("Fichier uploadé avec succès !");
    } else {
      setUploadMessage("Erreur lors de l'upload.");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen gap-4">
      <ThemeToggle />
      <h1 className="text-3xl font-bold">Bienvenue, {user} !</h1>

      {/* Composant d'upload */}
      <Card>
        <CardContent className="p-6 space-y-4">
          <div className="border-2 border-dashed border-gray-200 rounded-lg flex flex-col gap-1 p-6 items-center">
            <FileIcon className="w-12 h-12" />
            <span className="text-sm font-medium text-gray-500">Drag and drop a file or click to browse</span>
            <span className="text-xs text-gray-500">PDF, image, video, or audio</span>
          </div>
          <div className="space-y-2 text-sm">
            <Label htmlFor="file" className="text-sm font-medium">
              File
            </Label>
            <Input id="file" type="file" placeholder="File" accept="image/*" onChange={handleFileChange} />
          </div>
        </CardContent>
        <CardFooter>
          <Button size="lg" onClick={handleUpload} disabled={!selectedFile}>
            Upload
          </Button>
        </CardFooter>
      </Card>
      {uploadMessage && <p className="text-sm text-gray-700">{uploadMessage}</p>}

      {/* Bouton de déconnexion */}
      <button onClick={handleLogout} className="mt-4 px-4 py-2 bg-red-500 text-white rounded">
        Déconnexion
      </button>
    </div>
  );
}
