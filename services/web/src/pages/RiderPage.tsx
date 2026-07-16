import { Link } from 'react-router-dom';
import '../App.css';
import MapView from '../components/MapView';
import { useTrip } from '../hooks/useTrip';
import { useRiderWebSocket } from '../hooks/useRiderWebSocket';
import FareList from '../components/FareList';
import TripSummaryBar from '../components/TripSummaryBar';
import StatusBadge from '../components/StatusBadge';
import Button from '../components/Button';

// Generate a random ID for this session
const RIDER_ID = 'rider_' + Math.random().toString(36).substr(2, 9);

export function RiderPage() {
  const { 
    pickup, 
    destination, 
    route, 
    tripPreview, 
    selectedFare, 
    isLoading,
    error,
    handleMapClick, 
    selectFare 
  } = useTrip();
  
  const { isConnected, sendRideRequest } = useRiderWebSocket(RIDER_ID);

  const handleRequestRide = () => {
    if (pickup && destination && selectedFare) {
      sendRideRequest({
        pickup,
        destination,
        packageSlug: selectedFare.packageSlug,
      });
      // Simple visual feedback for now
      const btn = document.getElementById('request-btn');
      if (btn) {
        btn.innerText = 'Request Sent!';
        setTimeout(() => (btn.innerText = 'Request Ride'), 2000);
      }
    }
  };

  const renderSidebarContent = () => {
    if (!pickup) {
      return (
        <div className="sidebar-prompt">
          <div className="sidebar-prompt-icon">
            <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2} width={24}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              <path strokeLinecap="round" strokeLinejoin="round" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <p className="sidebar-prompt-text">Tap anywhere on the map to set your pickup location</p>
        </div>
      );
    }

    if (!destination) {
      return (
        <div className="sidebar-prompt">
          <div className="sidebar-prompt-icon">
            <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2} width={24}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M13 10V3L4 14h7v8l9-11h-7z" />
            </svg>
          </div>
          <p className="sidebar-prompt-text">Now tap to set your destination</p>
        </div>
      );
    }

    if (isLoading) {
      return (
        <div style={{ marginTop: '20px' }}>
          <div className="skeleton skeleton-summary" />
          <div className="skeleton skeleton-fare" />
          <div className="skeleton skeleton-fare" />
          <div className="skeleton skeleton-fare" />
          <div className="skeleton skeleton-fare" />
        </div>
      );
    }

    if (error) {
      return (
        <div className="sidebar-prompt" style={{ color: 'var(--color-danger)' }}>
          <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2} width={32} style={{ marginBottom: 12 }}>
            <path strokeLinecap="round" strokeLinejoin="round" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p>{error}</p>
        </div>
      );
    }

    if (tripPreview) {
      return (
        <div style={{ marginTop: '20px' }}>
          <TripSummaryBar 
            distance={tripPreview.route.distance} 
            duration={tripPreview.route.duration} 
          />
          <FareList 
            fares={tripPreview.rideFares} 
            selectedFare={selectedFare}
            onSelectFare={selectFare}
          />
        </div>
      );
    }

    return null;
  };

  return (
    <div className="app-container">
      <div className="sidebar">
        <div className="sidebar-inner">
          <div className="sidebar-header">
            <h2>Book a Ride</h2>
            <StatusBadge connected={isConnected} />
          </div>
          
          {renderSidebarContent()}

          <div className="sidebar-footer">
            {tripPreview && (
              <Button 
                id="request-btn"
                variant="primary" 
                full 
                onClick={handleRequestRide}
                disabled={!selectedFare || !isConnected}
              >
                {!isConnected 
                  ? 'Connecting...' 
                  : !selectedFare 
                    ? 'Select a ride' 
                    : `Request ${selectedFare.packageSlug} • ₹${selectedFare.totalPrice.toFixed(0)}`}
              </Button>
            )}
            <Link to="/" style={{ textDecoration: 'none' }}>
              <Button variant="ghost" full>Back to Home</Button>
            </Link>
          </div>
        </div>
      </div>

      <div className="map-container">
        <MapView 
          pickup={pickup} 
          destination={destination} 
          route={route} 
          onMapClick={(lat, lng) => handleMapClick({ latitude: lat, longitude: lng })} 
        />
      </div>
    </div>
  );
}
