import mongoose from 'mongoose';
import { NewsSchema, UserSchema } from './schemas.js';

const News = mongoose.model('News', NewsSchema);
const User = mongoose.model('User', UserSchema);

export { News, User };
