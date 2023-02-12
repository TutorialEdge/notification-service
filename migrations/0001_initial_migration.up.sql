-- Subscribers represent a subscriber to a given list within
-- a our system.
-- We default to is_subscribed on creation
CREATE TABLE subscribers (
    subscriber_id UUID primary key,
    email text NOT NULL,
    is_subscribed BOOL DEFAULT 't'
);

-- Notifications represent a specific type of notification within
-- the system. This could be a `WelcomeNotification` that we send out
-- whenever a new customer joins, or a `MonthlyUpdate` that we send out 
-- on a successful monthly payment 
CREATE TABLE notifications (
    notification_id UUID primary key,
    notification_name text,
    html text
);

-- List represents different subscriber groups, for example `Premium` which would
-- contain all subscribers that have a premium subscription for the site.
CREATE TABLE list (
    list_id UUID primary key,
    list_name text NOT NULL
);

 
