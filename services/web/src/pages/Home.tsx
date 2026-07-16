import { Link } from 'react-router-dom';
import '../App.css'; 

export function Home() {
  return (
    <div className="home">
      <h1 className="home-brand">GoRide</h1>
      <p className="home-subtitle">Your ride, a tap away</p>
      
      <div className="home-cards">
        <Link to="/rider" className="home-card">
          <div className="home-card-icon">
            <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </div>
          <h2>Book a Ride</h2>
          <p>Get a price and go</p>
        </Link>
        
        <Link to="/driver" className="home-card">
          <div className="home-card-icon">
            <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
            </svg>
          </div>
          <h2>Drive</h2>
          <p>Accept ride requests</p>
        </Link>
      </div>
    </div>
  );
}
