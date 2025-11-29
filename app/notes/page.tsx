"use client";

import React, { useState, useEffect } from "react";

type Note = {
  id: number;
  title: string;
  body: string;
  color: string; // pilihan warna card bebas (HEX, RGB, etc)
  date: string;
  time: string;
};

export default function NotesPage(): React.JSX.Element {
  const [notes, setNotes] = useState<Note[]>([]);
  const [editing, setEditing] = useState<Note | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem("token");

    fetch("/api/notes", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => res.json())
      .then(setNotes);
  }, []);

  function openNewNote() {
    setEditing({
      id: Date.now(),
      title: "",
      body: "",
      color: "#FFF59D",
      date: "",
      time: "",
    });
    setIsModalOpen(true);
  }

  function openEditNote(note: Note) {
    // clone to avoid uncontrolled updates
    setEditing({ ...note });
    setIsModalOpen(true);
  }

  const handleAdd = async () => {
    if (!editing) return;
    
    const token = localStorage.getItem("token");

    await fetch("/api/notes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ title: editing.title, content: editing.body }),
    });

    location.reload();
  };

  function saveNote(note: Note) {
    // Check if it's a new note (no existing id in database)
    const isNewNote = !notes.find((n) => n.id === note.id);
    
    if (isNewNote) {
      handleAdd();
      return;
    }

    // For existing notes, update locally (or you can add update API call here)
    setNotes((prev) => {
      const exists = prev.find((n) => n.id === note.id);
      if (exists) return prev.map((n) => (n.id === note.id ? note : n));
      return [note, ...prev];
    });

    setIsModalOpen(false);
    setEditing(null);
  }

  function deleteNote(id: number) {
    setNotes((prev) => prev.filter((n) => n.id !== id));
  }

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-7xl mx-auto">
        
        {/* HEADER */}
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-3xl font-semibold text-black">My Notes</h1>

          <button
            onClick={openNewNote}
            className="px-4 py-2 rounded-lg bg-black text-white text-sm"
          >
            New Note
          </button>
        </div>

        {/* Tabs visual */}
        <div className="flex gap-6 text-sm text-gray-500 mb-4">
          <button className="border-b-2 border-black pb-2">Today</button>
          <button>This Week</button>
          <button>This Month</button>
        </div>

        {/* notes grid */}
        <div className="grid grid-cols-3 gap-6">
          {notes.map((note) => (
            <article
              key={note.id}
              style={{ backgroundColor: note.color }}
              className="p-6 rounded-2xl shadow text-gray-900"
            >
              <div className="flex justify-between items-start mb-3">
                <span className="text-xs">
                  {note.date
                    ? new Date(note.date).toLocaleDateString()
                    : "No date"}
                </span>

                <div className="flex items-center gap-2">
                  <button
                    onClick={() => openEditNote(note)}
                    aria-label="edit"
                    className="p-1 rounded-md hover:bg-white/30"
                  >
                    ✎
                  </button>

                  <button
                    onClick={() => deleteNote(note.id)}
                    aria-label="delete"
                    className="p-1 rounded-md hover:bg-white/30"
                  >
                    🗑
                  </button>
                </div>
              </div>

              <h2 className="text-lg font-semibold mb-2">{note.title}</h2>

              <p className="text-sm leading-relaxed mb-6">{note.body}</p>

              <div className="text-xs text-gray-700 flex items-center gap-2">
                🕒 {note.time ? note.time : "No time"}
              </div>
            </article>
          ))}
        </div>

        {/* MODAL */}
        {isModalOpen && editing && (
          <div className="fixed inset-0 z-40 flex items-center justify-center">
            <div
              className="absolute inset-0 bg-black/40"
              onClick={() => setIsModalOpen(false)}
            />

            <div className="relative bg-white rounded-xl p-6 w-[520px] shadow-lg z-50">
              <h3 className="text-lg font-semibold mb-4 text-black">
                {editing.title ? "Edit Note" : "New Note"}
              </h3>

              {/* color preview bar */}
              <div
                className="w-full h-3 rounded mb-4"
                style={{ backgroundColor: editing.color }}
              />

              {/* Title */}
              <label className="block text-sm mb-2 text-gray-500">Title</label>
              <input
                className="w-full border rounded-md px-3 py-2 mb-3 text-gray-700"
                value={editing.title}
                onChange={(e) =>
                  setEditing({ ...editing, title: e.target.value })
                }
              />

              {/* Body */}
              <label className="block text-sm mb-2 text-gray-500">Body</label>
              <textarea
                className="w-full border rounded-md px-3 py-2 mb-3 h-28 text-gray-700"
                value={editing.body}
                onChange={(e) =>
                  setEditing({ ...editing, body: e.target.value })
                }
              />

              {/* DATE + TIME */}
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <label className="block text-sm text-gray-500 mb-1">
                    Date
                  </label>
                  <input
                    type="date"
                    value={editing.date}
                    onChange={(e) =>
                      setEditing({ ...editing, date: e.target.value })
                    }
                    className="w-full border rounded-md px-3 py-2 text-gray-700"
                  />
                </div>

                <div>
                  <label className="block text-sm text-gray-500 mb-1">
                    Time
                  </label>
                  <input
                    type="time"
                    value={editing.time}
                    onChange={(e) =>
                      setEditing({ ...editing, time: e.target.value })
                    }
                    className="w-full border rounded-md px-3 py-2 text-gray-700"
                  />
                </div>
              </div>

              {/* COLOR PICKER */}
              <div className="mb-4">
                <label className="block text-sm mb-2 text-gray-500">
                  Color
                </label>
                <input
                  type="color"
                  value={editing.color}
                  onChange={(e) =>
                    setEditing({ ...editing, color: e.target.value })
                  }
                  className="w-16 h-10 cursor-pointer"
                />
              </div>

              {/* BUTTONS */}
              <div className="flex items-center justify-end mt-4 gap-2">
                <button
                  onClick={() => {
                    setIsModalOpen(false);
                    setEditing(null);
                  }}
                  className="px-4 py-2 rounded-md border"
                >
                  Cancel
                </button>

                <button
                  onClick={() => {
                    if (!editing.title) return alert("Title required");
                    saveNote(editing);
                  }}
                  className="px-4 py-2 rounded-md bg-black text-white"
                >
                  Save
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
