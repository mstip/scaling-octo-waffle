import NewsCard from './NewsCard';

import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import axios from 'axios';

export default function News() {
  const [news, setNews] = useState([]);

  useEffect(async () => {
    const newsResponse = await axios.get(
      'https://3000-blush-sheep-06602up8.ws-eu17.gitpod.io/news'
    );
    setNews(newsResponse.data.news);
  }, []);

  return (
    <div>
      <h3>News (Latest)</h3>
      {news.map((item, index) => (
        <Link key={index} to={{ pathname: '/news/' + item._id }}>
          <NewsCard title={item.title} date={item.date} votes={item.votes} />
        </Link>
      ))}
    </div>
  );
}
