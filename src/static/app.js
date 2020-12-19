export async function fetchMessage() {
  try {
    const response = await fetch("/api/message");
    return await response.text();
  } catch (error) {
    return error.message;
  }
}
