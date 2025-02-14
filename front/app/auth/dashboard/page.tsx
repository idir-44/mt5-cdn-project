"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Table, TableHeader, TableRow, TableCell, TableBody } from "@/components/ui/table";
import { LucideFile as FileIcon, Folder as FolderIcon, Download, Upload } from "lucide-react";

interface FileItem {
  id: string;
  folderPath?: string;
  filename: string;
  filepath: string;
}

export default function Dashboard() {
  const [user, setUser] = useState<string | null>(null);
  const [files, setFiles] = useState<FileItem[]>([]);
  const [selectedFiles, setSelectedFiles] = useState<FileList | null>(null);
  const [uploadMessage, setUploadMessage] = useState<string | null>(null);
  const [downloadHistory, setDownloadHistory] = useState<string[]>([]);
  const [folder, setFolder] = useState<string>("");
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch("http://localhost:80/v1/me", {
          method: "GET",
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error("Échec de la récupération de l'utilisateur");
        }

        const data = await response.json();

        if (!data) {
          router.push("/auth/login");
        } else {
          setUser(data.email);
          await fetchFiles();
        }
      } catch (error) {
        console.error("Erreur lors de la récupération de l'utilisateur :", error);
        router.push("/auth/login"); // Rediriger en cas d'échec
      }
    };

    fetchUser();
  }, [router]);

  const fetchFiles = async () => {
    const response = await fetch(`http://localhost:80/v1/files?folder=${folder}`, {
      method: "GET",
      credentials: "include",
    });
    const data = await response.json();
    setFiles(data);
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      setSelectedFiles(event.target.files);
    }
  };

  const handleUpload = async () => {
    if (!selectedFiles) {
      setUploadMessage("Veuillez sélectionner des fichiers.");
      return;
    }
    console.log("Cookies envoyés :", document.cookie);

    const formData = new FormData();
    Array.from(selectedFiles).forEach((file) => formData.append("files", file));
    formData.append("folder", folder);

    const response = await fetch("http://localhost:80/v1/upload", {
      method: "POST",
      body: formData,
      credentials: "include",
    });

    if (response.ok) {
      setUploadMessage("Fichiers uploadés avec succès !");
      fetchFiles();
    } else {
      setUploadMessage("Erreur lors de l'upload.");
    }
  };

  const handleDownload = async (filepath: string) => {
    const response = await fetch(`http://localhost:80/v1/download?id=${filepath}`, {
      method: "GET",
      credentials: "include",
    });

    if (response.ok) {
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", filepath.split("/").pop() || "file");
      document.body.appendChild(link);
      link.click();
      setDownloadHistory([...downloadHistory, filepath]);
    }
  };

  const handleDownloadFolder = async () => {
    const response = await fetch(`http://localhost:8080/v1/download-folder?folder=${folder}`, {
      method: "GET",
      credentials: "include",
    });

    if (response.ok) {
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", `${folder}.zip`);
      document.body.appendChild(link);
      link.click();
    }
  };

  return (
      <div className="flex flex-col items-center justify-center min-h-screen gap-4">
        <h1 className="text-3xl font-bold">Bienvenue, {user} !</h1>

        {/* Upload Section */}
        <Card>
          <CardContent className="p-6 space-y-4">
            <Label htmlFor="folder">Dossier</Label>
            <Input
                id="folder"
                type="text"
                value={folder}
                onChange={(e) => setFolder(e.target.value)}
                placeholder="Nom du dossier"
            />

            <Label htmlFor="file">Fichiers</Label>
            <Input id="file" type="file" multiple onChange={handleFileChange} />
          </CardContent>
          <CardFooter>
            <Button size="lg" onClick={handleUpload} disabled={!selectedFiles}>
              <Upload className="mr-2" /> Upload
            </Button>
          </CardFooter>
        </Card>
        {uploadMessage && <p className="text-sm text-gray-700">{uploadMessage}</p>}

        {/* File List */}
        <h2 className="text-xl font-semibold mt-6">Fichiers & Dossiers</h2>
        <Table>
          <TableHeader>
            <TableRow>
              <TableCell>Nom</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHeader>
          <TableBody>
            {files?.map((file, index) => (
                <TableRow key={index}>
                  <TableCell>
                    {file.folderPath ? <FolderIcon className="inline-block mr-2" /> : <FileIcon className="inline-block mr-2" />}
                    {file.filename}
                  </TableCell>
                  <TableCell>
                    {!file.folderPath ? (
                        <Button variant="ghost" onClick={() => handleDownload(file.id)}>
                          <Download className="mr-2" /> Télécharger
                        </Button>
                    ) : (
                        <Button variant="ghost" onClick={handleDownloadFolder}>
                          <Download className="mr-2" /> Exporter
                        </Button>
                    )}
                  </TableCell>
                </TableRow>
            ))}
          </TableBody>
        </Table>

        {/* Download History */}
        <h2 className="text-xl font-semibold mt-6">Historique des téléchargements</h2>
        <ul>
          {downloadHistory.map((item, index) => (
              <li key={index}>{item}</li>
          ))}
        </ul>
      </div>
  );
}