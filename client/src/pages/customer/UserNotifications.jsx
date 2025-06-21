import {
  useBrowserNotificationsQuery,
  useMarkAllNotificationsAsRead,
} from "@/hooks/useNotification";
import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/Badge";
import { formatDistanceToNow } from "date-fns";
import { CheckCheck, MailWarning } from "lucide-react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/Tabs";

const UserNotifications = () => {
  const [tab, setTab] = useState("unread");
  const [hasMarkedRead, setHasMarkedRead] = useState(false);
  const { mutate: markAllAsRead } = useMarkAllNotificationsAsRead();
  const { data, isLoading, isError, refetch } = useBrowserNotificationsQuery();

  const notifications = data || [];
  const unread = notifications.filter((n) => !n.isRead);
  const read = notifications.filter((n) => n.isRead);

  useEffect(() => {
    if (tab === "unread" && unread.length > 0 && !hasMarkedRead) {
      markAllAsRead();
      setHasMarkedRead(true);
    }
  }, [tab, unread.length, hasMarkedRead, markAllAsRead]);

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const renderNotifications = (list) =>
    list.length === 0 ? (
      <div className="flex flex-col items-center text-muted-foreground py-10 border border-dashed rounded-xl space-y-2 mt-2">
        <MailWarning className="w-10 h-10" />
        <p className="text-sm">No {tab} notifications</p>
      </div>
    ) : (
      <div className="space-y-4">
        {list.map((notif) => (
          <div
            key={notif.id}
            className={`rounded-lg p-4 border ${
              notif.isRead ? "bg-muted/50" : "bg-background"
            } transition`}
          >
            <div className="flex items-start justify-between">
              <div className="space-y-1 max-w-[85%]">
                <h4 className="font-semibold">{notif.title}</h4>
                <p className="text-sm text-muted-foreground">{notif.message}</p>
                <p className="text-xs text-muted-foreground italic">
                  {formatDistanceToNow(new Date(notif.createdAt), {
                    addSuffix: true,
                  })}
                </p>
              </div>
              <div className="flex items-center gap-2">
                {!notif.isRead && <Badge variant="default">New</Badge>}
                <CheckCheck
                  className={`w-5 h-5 ${
                    notif.isRead ? "text-muted" : "text-primary"
                  }`}
                />
              </div>
            </div>
          </div>
        ))}
      </div>
    );

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Inbox"
        description="See all your recent notifications"
      />
      <Tabs defaultValue="unread" onValueChange={(val) => setTab(val)}>
        <TabsList className="mb-4">
          <TabsTrigger value="unread">Unread ({unread.length})</TabsTrigger>
          <TabsTrigger value="read">Read ({read.length})</TabsTrigger>
        </TabsList>

        <TabsContent value="unread">{renderNotifications(unread)}</TabsContent>
        <TabsContent value="read">{renderNotifications(read)}</TabsContent>
      </Tabs>
    </section>
  );
};

export default UserNotifications;
