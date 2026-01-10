import { apiClient } from "@/src/lib/api/client";
import { NextResponse } from "next/server";

export async function POST (request: Request) {
    const { email, password } = await request.json();
    const response = await apiClient.post('/v1/users/auth/login', { email, password });
    return NextResponse.json(response.data);
}