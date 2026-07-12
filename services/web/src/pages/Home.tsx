import { Link } from 'react-router-dom';
import '../App.css'; // Reusing some base styles

export function Home() {
  return (
    <div className="home-container" style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100vh', backgroundColor: '#f0f2f5' }}>
      <h1>Welcome to GoRide Platform</h1>
      <p style={{ marginBottom: '2rem', color: '#666' }}>Select an application to launch</p>
      
      <div style={{ display: 'flex', gap: '2rem' }}>
        <Link to="/rider" className="card" style={{ textDecoration: 'none', color: 'inherit', padding: '2rem', textAlign: 'center', minWidth: '200px', cursor: 'pointer' }}>
          <div style={{ fontSize: '3rem', marginBottom: '1rem' }}>🧑‍s</div>
          <h2>Rider App</h2>
          <p>Book a ride</p>
        </Link>
        
        <Link to="/driver" className="card" style={{ textDecoration: 'none', color: 'inherit', padding: '2rem', textAlign: 'center', minWidth: '200px', cursor: 'pointer' }}>
          <div style={{ fontSize: '3rem', marginBottom: '1rem' }}>🚗</div>
          <h2>Driver App</h2>
          <p>Accept rides</p>
        </Link>
      </div>
    </div>
  );
}
