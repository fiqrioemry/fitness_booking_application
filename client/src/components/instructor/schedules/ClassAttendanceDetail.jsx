import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetDescription,
} from "@/components/ui/Sheet";
import { Badge } from "@/components/ui/Badge";
import { useNavigate, useParams } from "react-router-dom";
import { useScheduleAttendanceQuery } from "@/hooks/useSchedules";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/Avatar";
import { ClassAttendanceSkeleton } from "@/components/loading/ClassAttendanceSkeleton";

export const ClassAttendanceDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { data, isLoading } = useScheduleAttendanceQuery(id);

  return (
    <Sheet open={true} onOpenChange={() => navigate(-1)}>
      <SheetContent side="right" className="max-w-xl w-full">
        <SheetHeader>
          <SheetTitle>Class Schedule Detail</SheetTitle>
          <SheetDescription>Your Student attendance info</SheetDescription>
        </SheetHeader>

        {isLoading ? (
          <ClassAttendanceSkeleton />
        ) : (
          <div className="mt-6 space-y-4">
            {data.map((user, index) => (
              <div
                key={index}
                className="flex items-center gap-4 border border-muted rounded-xl px-4 py-3"
              >
                <Avatar className="w-10 h-10">
                  <AvatarImage src={user?.avatar} />
                  <AvatarFallback>
                    {user.fullname
                      .split(" ")
                      .map((n) => n[0])
                      .join("")
                      .toUpperCase()}
                  </AvatarFallback>
                </Avatar>

                <div className="flex-1">
                  <p className="text-sm font-medium">{user.fullname}</p>
                  <p className="text-xs text-muted-foreground">{user.email}</p>
                </div>

                <Badge variant="secondary" className="text-xs capitalize">
                  {user.status}
                </Badge>
              </div>
            ))}
          </div>
        )}
      </SheetContent>
    </Sheet>
  );
};
