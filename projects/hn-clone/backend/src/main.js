import Express from 'express';
import cors from 'cors';
import mongoose from 'mongoose';
import { v4 as uuidv4 } from 'uuid';

import { News } from './models.js';

const app = Express();
app.use(cors());
app.use(Express.json());
const port = 3000;

app.get('/', async (req, res) => {
  res.send('hi!');
});

app.get('/news', async (req, res) => {
  res.send({ news: await News.find() });
});

app.get('/news/:id', async (req, res) => {
  const newsId = req.params.id;
  res.send(await News.findById(newsId));
});

app.post('/user/login', async (req, res) => {
  console.log(req.body);
  res.send({token: uuidv4()});
});

async function main() {
  await mongoose.connect('mongodb://localhost:27017/hn-clone');
  app.listen(port, () => {
    console.log(`running at http://localhost:${port}`);
  });
}

await main();
