import News from './views/news/News';
import NewsDetail from './views/detail/NewsDetail';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Login from './views/login/login';

function App() {
  return (
    <Router>
      <div>
        <Link to='/'>Start</Link>
        <span> | </span>
        <Link to='/'>Post News</Link>
        <span> | </span>
        <Link to='/'>Latest</Link>
        <span> | </span>
        <Link to='/'>Hot</Link>
        <span> | </span>
        <Link to='/login'>Login</Link>
        <hr className='border-t-2 border-red-300' />
        <Switch>
          <Route path='/news/:newsId'>
            <NewsDetail />
          </Route>
          <Route path='/login'>
            <Login />
          </Route>
          <Route path='/'>
            <News />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
