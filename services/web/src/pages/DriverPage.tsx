import { useState } from 'react';
import '../App.css';
import MapView from '../components/MapView';
import { useDriverWebSocket } from '../hooks/useDriverWebSocket';
import { Link } from 'react-router-dom';

const DRIVER_ID = 'driver_' + Math.random().toString(36).substr(2, 9);

export function DriverPage() {
  const [packageSlug, setPackageSlug] = useState('premium');
  const [shouldConnect, setShouldConnect] = useState(false);
  
  // Only connect if the user hits "Connect"
  const activeSlug = shouldConnect ? packageSlug : '';
  const { isConnected, driverInfo, availableRides } = useDriverWebSocket(
    shouldConnect ? DRIVER_ID : '', 
    activeSlug
  );

  return (
    <div className="app-container">
      <div className="map-container">
        {/* Simple map centered on a default location, can be enhanced with driver location later */}
        <MapView 
          pickup={null} 
          destination={null} 
          route={null} 
          onMapClick={() => {}} 
        />
      </div>
      
      <div className="sidebar">
        <div className="card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <h2>Driver Dashboard</h2>
            <div title="WebSocket Status">
               {isConnected ? '🟢 Connected' : '🔴 Disconnected'}
            </div>
          </div>

          {!isConnected ? (
            <div style={{ marginTop: '2rem' }}>
              <label style={{ display: 'block', marginBottom: '0.5rem', fontWeight: 'bold' }}>
                Select Package:
              </label>
              <select 
                value={packageSlug}
                onChange={(e) => setPackageSlug(e.target.value)}
                style={{ width: '100%', padding: '0.8rem', borderRadius: '8px', marginBottom: '1rem', border: '1px solid #ccc' }}
              >
                <option value="premium">Premium 👑</option>
                <option value="suv">SUV 🚐</option>
                <option value="sedan">Sedan 🚗</option>
                <option value="van">Van 🚌</option>
              </select>

              <button 
                className="back-button" 
                style={{ width: '100%', backgroundColor: '#000', color: '#fff', marginTop: '1rem' }}
                onClick={() => setShouldConnect(true)}
              >
                Go Online
              </button>
            </div>
          ) : (
            <>
              {driverInfo && (
                <div style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  gap: '1rem', 
                  marginTop: '1.5rem', 
                  padding: '1rem', 
                  backgroundColor: '#f8f9fa', 
                  borderRadius: '12px' 
                }}>
                  <div style={{ 
                    width: '60px', 
                    height: '60px', 
                    borderRadius: '50%', 
                    backgroundColor: '#ccc',
                    backgroundImage: `url(${driverInfo.profilePicture})`,
                    backgroundSize: 'cover',
                    backgroundPosition: 'center'
                  }} />
                  <div>
                    <h3 style={{ margin: 0 }}>{driverInfo.name}</h3>
                    <p style={{ margin: '0.2rem 0', color: '#666' }}>{driverInfo.carNumber} • {driverInfo.packageSlug}</p>
                  </div>
                </div>
              )}

              <div style={{ marginTop: '2rem' }}>
                <h3>Available Rides</h3>
                {availableRides.length === 0 ? (
                  <p style={{ color: '#666', fontStyle: 'italic', padding: '1rem 0' }}>
                    Listening for new ride requests...
                  </p>
                ) : (
                  <div className="ride-list" style={{ marginTop: '1rem' }}>
                    {availableRides.map((ride, idx) => (
                      <div key={idx} className="ride-item" style={{ padding: '1rem', border: '1px solid #eaeaea', borderRadius: '12px' }}>
                        <div><strong>Pickup:</strong> {ride.pickup?.latitude}, {ride.pickup?.longitude}</div>
                        <div><strong>Dest:</strong> {ride.destination?.latitude}, {ride.destination?.longitude}</div>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              <button 
                className="back-button" 
                style={{ width: '100%', marginTop: '1rem' }}
                onClick={() => setShouldConnect(false)}
              >
                Go Offline
              </button>
            </>
          )}

          <div style={{ marginTop: '1rem' }}>
            <Link to="/" className="back-button" style={{ display: 'block', textDecoration: 'none', textAlign: 'center' }}>
              Home
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
