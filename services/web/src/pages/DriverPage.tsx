import { useState } from 'react';
import { Link } from 'react-router-dom';
import '../App.css';
import MapView from '../components/MapView';
import { useDriverWebSocket } from '../hooks/useDriverWebSocket';
import StatusBadge from '../components/StatusBadge';
import Button from '../components/Button';

const DRIVER_ID = 'driver_' + Math.random().toString(36).substr(2, 9);

export function DriverPage() {
  const [packageSlug, setPackageSlug] = useState('sedan');
  const [shouldConnect, setShouldConnect] = useState(false);
  
  const activeSlug = shouldConnect ? packageSlug : '';
  const { isConnected, driverInfo, availableRides } = useDriverWebSocket(
    shouldConnect ? DRIVER_ID : '', 
    activeSlug
  );

  return (
    <div className="app-container">
      <div className="sidebar">
        <div className="sidebar-inner">
          <div className="sidebar-header">
            <h2>Driver App</h2>
            <StatusBadge connected={isConnected} />
          </div>

          {!isConnected ? (
            <div className="driver-section">
              <div className="select-wrapper">
                <label>Vehicle Class</label>
                <select 
                  value={packageSlug}
                  onChange={(e) => setPackageSlug(e.target.value)}
                >
                  <option value="sedan">Sedan (Standard)</option>
                  <option value="suv">SUV (6 Seats)</option>
                  <option value="van">Van (8+ Seats)</option>
                  <option value="luxury">Luxury (Premium)</option>
                </select>
              </div>

              <Button 
                variant="primary" 
                full 
                onClick={() => setShouldConnect(true)}
                style={{ marginTop: 8 }}
              >
                Go Online
              </Button>
            </div>
          ) : (
            <>
              {driverInfo && (
                <div className="driver-profile">
                  <div 
                    className="driver-avatar"
                    style={{ backgroundImage: `url(${driverInfo.profilePicture || 'https://api.dicebear.com/7.x/avataaars/svg?seed=' + driverInfo.name})` }} 
                  />
                  <div>
                    <h3 className="driver-name">{driverInfo.name}</h3>
                    <p className="driver-meta">{driverInfo.carNumber} • {driverInfo.packageSlug.toUpperCase()}</p>
                  </div>
                </div>
              )}

              <div className="driver-section">
                <h3>Incoming Requests</h3>
                {availableRides.length === 0 ? (
                  <div className="driver-rides-empty">
                    Waiting for nearby requests...
                  </div>
                ) : (
                  <div className="fare-list" style={{ marginTop: 12 }}>
                    {availableRides.map((ride, idx) => (
                      <div key={idx} className="driver-ride-item">
                        <div style={{ marginBottom: 4 }}>
                          <strong>Pickup:</strong> {ride.pickup?.latitude.toFixed(4)}, {ride.pickup?.longitude.toFixed(4)}
                        </div>
                        <div>
                          <strong>Drop:</strong> {ride.destination?.latitude.toFixed(4)}, {ride.destination?.longitude.toFixed(4)}
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              <div style={{ marginTop: 'auto', paddingTop: 24 }}>
                <Button 
                  variant="secondary" 
                  full 
                  onClick={() => setShouldConnect(false)}
                >
                  Go Offline
                </Button>
              </div>
            </>
          )}

          <div className="sidebar-footer">
            <Link to="/" style={{ textDecoration: 'none' }}>
              <Button variant="ghost" full>Back to Home</Button>
            </Link>
          </div>
        </div>
      </div>
      
      <div className="map-container">
        <MapView 
          pickup={null} 
          destination={null} 
          route={null} 
          onMapClick={() => {}} 
        />
      </div>
    </div>
  );
}
