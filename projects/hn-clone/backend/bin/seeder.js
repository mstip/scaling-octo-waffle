import mongoose from 'mongoose';
import { News, User } from '../src/models.js';
import faker from 'faker';
import bcrypt from 'bcrypt';

async function seeder() {
  await mongoose.connect('mongodb://localhost:27017/hn-clone');

  const pwHash = await bcrypt.hash('qqqq', 10);
  for (let i = 0; i < 100; i++) {
    const user = await new User({
      name: faker.internet.userName(),
      email: faker.internet.exampleEmail(),
      passwordHash: pwHash
    }).save();

    const title = faker.lorem.words(3);
    await new News({
      title,
      url: faker.internet.url(),
      body: faker.lorem.text(),
      date: new Date(),
      votes: 1,
      author: user._id,
      comments: [
        {
          body: faker.lorem.text(),
          date: new Date(),
          votes: 3,
          comments: [
            {
              body: faker.lorem.text(),
              date: new Date(),
              votes: 3,
            },
            {
              body: faker.lorem.text(),
              date: new Date(),
              votes: 3,
              comments: [
                {
                  body: faker.lorem.text(),
                  date: new Date(),
                  votes: 3,
                },
                {
                  body: faker.lorem.text(),
                  date: new Date(),
                  votes: 3,
                },
              ],
            },
            {
              body: faker.lorem.text(),
              date: new Date(),
              votes: 3,
            },
          ],
        },
      ],
    }).save();

    console.log(`${user._id} - ${i} - ${title}`);
  }

  await mongoose.connection.close();
}

await seeder();
