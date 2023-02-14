-- Subscribers represent a subscriber to a given list within
-- a our system.
-- We default to is_subscribed on creation
CREATE TABLE subscribers (
    subscriber_id UUID primary key,
    email text NOT NULL,
    is_subscribed BOOL DEFAULT 't',
    created_at TIMESTAMP DEFAULT CURRENT_TIME,
    updated_at TIMESTAMP DEFAULT CURRENT_TIME,
    deleted_at TIMESTAMP
);

-- Notifications represent a specific type of notification within
-- the system. This could be a `WelcomeNotification` that we send out
-- whenever a new customer joins, or a `MonthlyUpdate` that we send out 
-- on a successful monthly payment 
CREATE TABLE notifications (
    notification_id UUID primary key,
    notification_name text,
    html text,
    created_at TIMESTAMP DEFAULT CURRENT_TIME,
    updated_at TIMESTAMP DEFAULT CURRENT_TIME,
    deleted_at TIMESTAMP
);

-- List represents different subscriber groups, for example `Premium` which would
-- contain all subscribers that have a premium subscription for the site.
CREATE TABLE list (
    list_id UUID primary key,
    list_name text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIME,
    updated_at TIMESTAMP DEFAULT CURRENT_TIME,
    deleted_at TIMESTAMP
);

 
