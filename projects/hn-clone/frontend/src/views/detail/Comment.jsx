import { Link } from 'react-router-dom';

export default function Comment({ body, author, votes, date, link }) {
  return (
    <div className='p-2 border-black border-b-2'>
      <p>{body}</p>
      <span className='text-sm font-light'>
        {date} by {author} with {votes} votes -
        <Link to={{ pathname: link + '/answer' }}>Answer</Link> -
        <Link to={{ pathname: link + '/upvote' }}>Upvote</Link> -
        <Link to={{ pathname: link + '/downvote' }}>Downvote</Link>
      </span>
    </div>
  );
}
