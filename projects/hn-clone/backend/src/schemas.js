import mongoose from 'mongoose';

const CommentsSchema = new mongoose.Schema({
  body: String,
  author: mongoose.ObjectId,
  date: Date,
  votes: Number,
  comments: [],
});

const NewsSchema = new mongoose.Schema({
  title: String,
  url: String,
  body: String,
  author: mongoose.ObjectId,
  date: Date,
  votes: Number,
  comments: [CommentsSchema],
});

const UserSchema = new mongoose.Schema({
  name: String,
  email: String,
  passwordHash: String,
  token: String
}); 


export { NewsSchema, CommentsSchema, UserSchema };
