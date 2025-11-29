"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { postJson } from "@/app/lib/api";

export default function SignUpPage() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const handleRegister = async () => {
    setError(""); // Clear previous errors
    
    if (!username || !email || !password) {
      setError("Please fill in all fields");
      return;
    }

    const res = await postJson("/api/auth/register", { username, email, password });

    if (res.ok) {
      alert("Registration successful!");
      router.push("/login");
    } else {
      const errorMsg = typeof res.data === 'object' && res.data.error 
        ? res.data.error 
        : "Registration failed";
      setError(errorMsg);
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-[#f7f8fc] px-4">
      <h1 className="text-3xl font-semibold text-[#5a5eb9] mb-1">Sign up</h1>
      <p className="text-gray-500 mb-6">Join the community today!</p>

      <button className="flex items-center gap-2 w-full max-w-md bg-white py-3 rounded-full shadow text-gray-700 justify-center mb-4">
        <span className="text-xl">G</span>
        Use Google account
      </button>

      <div className="text-gray-400 mb-4">or</div>

      {error && (
        <div className="w-full max-w-md bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-4">
          {error}
        </div>
      )}

      <form 
        className="w-full max-w-md flex flex-col gap-4"
        onSubmit={(e) => {
          e.preventDefault();
          handleRegister();
        }}
      >
        <div className="flex flex-col w-full">
          <label className="text-gray-600 mb-1 font-bold">Email</label>
          <input
            type="email"
            className="border-b border-gray-300 bg-transparent focus:outline-none py-2 text-gray-500"
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>

        <div className="flex flex-col w-full">
          <label className="text-gray-600 mb-1 font-bold">Username</label>
          <input
            type="text"
            className="border-b border-gray-300 bg-transparent focus:outline-none py-2 text-gray-500"
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>

        <div className="flex flex-col w-full">
          <label className="text-gray-600 mb-1 font-bold">Password</label>
          <input
            type="password"
            className="border-b border-gray-300 bg-transparent focus:outline-none py-2 text-gray-500"
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <button
          type="submit"
          className="w-full bg-gradient-to-r from-[#7d8cf3] to-[#5a5eb9] text-white py-3 rounded-full mt-2"
        >
          Sign up
        </button>
      </form>

      <p className="text-gray-600 mt-6">
        Already a member? <Link href="/login" className="text-[#5a5eb9] hover:underline cursor-pointer">Sign in</Link>
      </p>
    </div>
  );
}

