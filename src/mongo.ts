import { MongoClient } from "https://deno.land/x/mongo@v0.23.0/mod.ts";
import { config } from "https://deno.land/x/dotenv@v2.0.0/mod.ts";

export async function insert(
  json: { [key: string]: string | number; date: string },
  collection: string,
): Promise<void> {
  const env = config();
  const client = new MongoClient();
  await client.connect(
    `mongodb://${env.MONGO_USER}:${env.MONGO_PASSWORD}@localhost:27017`,
  );
  const db = client.database(env.MONGO_DATABASE);
  const sensors = db.collection(collection);
  await sensors.insertOne(json);
}
