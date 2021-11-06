export default function Login() {
  return (
    <div className='p-2 mx-20 my-2 border-black border-2'>
      <h3 className=''>Login</h3>
      <form>
        <div>
          <label>Username</label>
          <input type='text' name='username' className='border-black border-2' />
        </div>
        <div>
          <label>Password</label>
          <input type='password' name='password' className='border-black border-2' />
        </div>
        <button type='submit' className='border-black border-2' >
          Login
        </button>
      </form>
    </div>
  );
}
