import {
  useNotificationSettingsQuery,
  useUpdateNotificationSetting,
} from "@/hooks/useNotification";
import { Switch } from "@/components/ui/Switch";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";

const groupedByTitle = (notifications) => {
  return notifications.reduce((acc, item) => {
    if (!acc[item.title]) acc[item.title] = [];
    acc[item.title].push(item);
    return acc;
  }, {});
};

const labelMap = {
  sms: "SMS",
  email: "Email",
  browser: "Browser notification",
};

const UserSettings = () => {
  const { data, isError, refetch, isLoading } = useNotificationSettingsQuery();

  const { mutate: updateSetting } = useUpdateNotificationSetting();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const notifications = data || [];
  const grouped = groupedByTitle(notifications);

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="Notification settings"
        description="Choose how you'd like to be notified"
      />
      {Object.entries(grouped).map(([title, list]) => (
        <div key={title} className="border-b pb-6 space-y-4">
          <div>
            <h3 className="font-semibold text-lg">{title}</h3>
            <p className="text-muted-foreground text-sm">
              {title === "New Promotion Available"
                ? "Stay up-to-date with our deals and promotions."
                : title === "Daily Class Reminder"
                ? "Get a reminder of your scheduled classes every day."
                : title === "Class Reminder"
                ? "Be notified 1 hour before your class starts."
                : title === "Booking Successful"
                ? "Get confirmation for every successful booking."
                : title === "Payment Successful"
                ? "Get notified after payment is complete."
                : "Receive updates from our system."}
            </p>
          </div>

          <div className="space-y-8">
            {list.map((item) => (
              <div
                key={`${item.typeId}-${item.channel}`}
                className="flex items-center justify-between"
              >
                <span className="text-sm">{labelMap[item.channel]}</span>
                <Switch
                  checked={item.enabled}
                  onCheckedChange={(val) =>
                    updateSetting({
                      typeId: item.typeId,
                      channel: item.channel,
                      enabled: val,
                    })
                  }
                  className="transition duration-200"
                />
              </div>
            ))}
          </div>
        </div>
      ))}
    </section>
  );
};

export default UserSettings;
