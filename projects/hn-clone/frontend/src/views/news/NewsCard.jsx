export default function NewsCard({ title, date, votes, author }) {
  return (
    <div className='p-2 mx-20 my-2 border-black border-2'>
      <div>{title}</div>
      <div>
        {date} from {author} with {votes} votes
      </div>
    </div>
  );
}
