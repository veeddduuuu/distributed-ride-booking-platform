interface TripSummaryBarProps {
  distance: number; // in meters
  duration: number; // in seconds
}

export default function TripSummaryBar({ distance, duration }: TripSummaryBarProps) {
  const km = (distance / 1000).toFixed(1);
  const minutes = Math.round(duration / 60);

  return (
    <div className="trip-summary">
      <div className="trip-summary-item">
        <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
          <path strokeLinecap="round" strokeLinejoin="round" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
        </svg>
        <span className="trip-summary-value">{km} km</span>
      </div>
      <div className="trip-summary-dot" />
      <div className="trip-summary-item">
        <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
          <path strokeLinecap="round" strokeLinejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span className="trip-summary-value">{minutes} min</span>
      </div>
    </div>
  );
}
