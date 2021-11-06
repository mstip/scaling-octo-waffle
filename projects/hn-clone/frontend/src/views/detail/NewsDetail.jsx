import { useEffect, useState } from 'react';
import { useParams } from 'react-router';
import axios from 'axios';
import Comment from './Comment';

export default function NewsDetail() {
  const { newsId } = useParams();

  const [news, setNews] = useState({});

  useEffect(async () => {
    const newsResponse = await axios.get(
      'https://3000-blush-sheep-06602up8.ws-eu17.gitpod.io/news/' + newsId
    );
    setNews(newsResponse.data);
  }, []);

  function createComments(data, level = 0) {
    if (!data?.comments?.length) {
      return null;
    }
    return data.comments?.map((item, index) => {
      const els = [
        <Comment
          key={`${level}-${index}`}
          body={item.body}
          author={item.author}
          date={item.date}
          votes={item.votes}
          link={`/news/${newsId}/${level}-${index}`}
        />,
      ];
      if (item?.comments?.length) {
        els.push(
          <div key={level} className='mx-20'>
            {createComments(item, level + 1)}
          </div>
        );
      }
      return els;
    });
  }

  return (
    <div>
      <div className='p-2 mx-20 my-2 border-black border-2'>
        <h1 className='text-xl font-bold'>{news.title}</h1>
        <a className='text-sm italic' href={news.url}>
          {news.url}
        </a>
        <div>{news.body}</div>
        <span className='text-sm font-light'>
          {news.date} by {news.author} with {news.votes} votes
        </span>
      </div>
      <div className='mx-20'>{createComments(news)}</div>
    </div>
  );
}
