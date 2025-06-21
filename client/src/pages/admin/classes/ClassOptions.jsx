import { useState } from "react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSelectOptions } from "@/hooks/useSelectOptions";
import { SectionTitle } from "@/components/header/SectionTitle";
import { AddOptions } from "@/components/admin/classes/AddOptions";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { UpdateOptions } from "@/components/admin/classes/UpdateOptions";
import { DeleteOptions } from "@/components/admin/classes/DeleteOptions";

const tabs = [
  { label: "Type", type: "type" },
  { label: "Level", type: "level" },
  { label: "Category", type: "category" },
  { label: "Subcategory", type: "subcategory" },
  { label: "Location", type: "location" },
];

const ClassOptions = () => {
  const [activeTab, setActiveTab] = useState(tabs[0].type);
  const {
    data: options = [],
    isLoading,
    isError,
    refetch,
  } = useSelectOptions(activeTab);

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const renderMap = (geoLocation) => {
    const [lat, lng] = geoLocation.split(",");
    const mapUrl = `https://maps.google.com/maps?q=${lat},${lng}&hl=es&z=14&output=embed`;
    return (
      <div className="rounded overflow-hidden border border-border">
        <iframe
          src={mapUrl}
          width="100%"
          height="375"
          loading="lazy"
          allowFullScreen
          title="location-map"
          className="rounded-md"
        ></iframe>
      </div>
    );
  };

  return (
    <section className="section">
      <SectionTitle
        title="Class Options"
        description="Manage all static data such as class types, categories, levels, and
          locations."
      />

      <div className="flex gap-2 border-b border-border overflow-x-auto mt-4">
        {tabs.map((tab) => (
          <button
            key={tab.type}
            onClick={() => setActiveTab(tab.type)}
            className={`px-4 py-2 text-sm font-medium whitespace-nowrap transition rounded-t-md ${
              activeTab === tab.type
                ? "bg-muted text-primary border border-border border-b-transparent"
                : "text-muted-foreground hover:text-primary"
            }`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      <div className="space-y-4 mt-6">
        <div className="flex justify-between items-center">
          <h3 className="text-lg font-semibold capitalize text-foreground">
            {activeTab} List
          </h3>
          <AddOptions activeTab={activeTab} />
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {options.map((item) => (
            <div
              key={item.id}
              className="flex flex-col justify-between gap-2 border border-border bg-background rounded-xl px-4 py-3 shadow-sm hover:shadow-md transition"
            >
              <div className="flex-1 space-y-1">
                <span className="font-medium text-foreground block truncate">
                  {item.name}
                </span>
                {activeTab === "location" && (
                  <>
                    <p className="text-sm text-muted-foreground truncate">
                      {item.address}
                    </p>
                    {item.geoLocation && renderMap(item.geoLocation)}
                  </>
                )}
              </div>
              <div className="flex justify-end gap-2 pt-2">
                <DeleteOptions option={item} activeTab={activeTab} />
                <UpdateOptions option={item} activeTab={activeTab} />
              </div>
            </div>
          ))}

          {options.length === 0 && (
            <p className="text-sm text-muted-foreground col-span-full">
              No data found.
            </p>
          )}
        </div>
      </div>
    </section>
  );
};

export default ClassOptions;
