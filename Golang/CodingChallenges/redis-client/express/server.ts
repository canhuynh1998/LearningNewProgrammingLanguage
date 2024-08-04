import express, { Express } from 'express';
import { Redis } from 'ioredis';

const app: Express = express();
const port = process.env.PORT || 3000;

const redis = new Redis(6380);

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', async (req, res) => {
  const value = await redis.get('hi')
  console.log(value);
  res.send({value:value});
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});