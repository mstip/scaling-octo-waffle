import dotenv from 'dotenv';
dotenv.config();

setInterval(() => console.log(`hello from server ${process.env.SERVER_ID}`), process.env.POLL_INTERVAL);
