import ModeToggle from "@/components/ui/theme-toogle";


export default function Header() {
    return ( 
        <header className={`flex justify-between items-center py-4 px-8`}>
            <div />
            <h1 className={`text-2xl`}>Web application programming</h1>
            <ModeToggle/>
        </header>
     );
}

